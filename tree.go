package tree

import (
	"encoding/binary"
	"fmt"

	"github.com/cyphrme/coze"
)

// See tree/README.md documentation.
//
// Recursive options are only populated if the currently branch level is not
// leaves and contains a deeper level.
//
// # Structure Variables
//
// Alg: Hashing algorithm used for creating the tree.
//
// Seed: "Private" seed.  Must be same size as digest.
//
// ID: "Public" seed identifier.  The digest of the seed is the identifier, and
// is not apart of Branch.
//
// BLS: "Branch Level Sizes"  Size for each level of a recursive tree.  Populate
// assumes a symmetrical tree.  The last level may be non-symmetrical if
// MaxTotalLeaves is set.  The hypothetical "Branch level" variable, which would
// be an integer representing how deep a branch is, is not stored as an explicit
// variable but is known by the parent from the values of BLS.  A "Branch"
// variable would be just the seed of a branch as named by the parent.
//
// Branches: The digest values of the branches based on root seed for this
// level. Each branch value is the seed for the respective children.  At the
// edge of the tree, Branches are termed "leaves" and have no further children
// calculated.  Branches does not include children branches.
//
// Skip: Optional starting position for the first level.  Populates tree
// starting at this position.
//
// PathCalc:  If true, Paths, PathsID, Leaves, and LeavesID are calculated.
//
// Paths: ("Private") The tree path (all branches above) for all nodes.  All
// leaves in a branch share the same Path. Both keys and values are seeds.
// Object of keys as string, values as []B64 of path, e.g.  "paths": {"77Jo...":
// ["0RJw..."],
//
// PathsID: ("Public") Like paths except uses ID's as the keys and values.
//
// Leaves: ("Private") A slice of leaves (their seeds) in this tree down.  Only
// calculated if LeafPathCalc.
//
// LeavesID: ("Public") Like Leaves except uses ID's as the keys and values.
//
// ## Stats and Internal Variables
//
// TotalLeaves: (Breaks one-way design pattern) Number of leaves currently
// populated in the whole tree, including "skip".  (Children trees
// `TreeTotalLeaves` will be parent tree's `TreeTotalLeaves` as well.) Leaves
// are the last branch with no children.
//
// MaxTotalLeaves: (Breaks one-way design pattern) Normally left empty.  Total
// leaves in the whole tree. (recursive data structure).  If TotalLeaves is
// empty, each branch is fully populated  based on the BLS.  If TotalLeaves is
// populated, the last populated branch may be unsymmetrical when the
// MaxTotalLeaves limit is reached.
//
// ## Children
//
// Children: Recursive tree.  Each digest in "Branch" is the key for the tree in
// "BranchTree" on a one to one basis.  Note that the "public" value for the
// branch is used for the generation of a child tree, and thus becomes "private"
// from that perspective.
type Tree struct {
	Alg  coze.HshAlg `json:"alg"`
	Seed coze.B64    `json:"seed"`
	ID   coze.B64    `json:"id"`
	BLS  []int       `json:"branch_level_sizes,omitempty"`

	Skip     int        `json:"skip,omitempty"`
	Branches []coze.B64 `json:"branches,omitempty"`

	PathCalc bool       `json:"path_calc,omitempty"`
	Paths    B64MapP    `json:"paths,omitempty"`
	PathsID  B64MapP    `json:"paths_id,omitempty"`
	Leaves   []coze.B64 `json:"leaves,omitempty"`
	LeavesID []coze.B64 `json:"leaves_id,omitempty"`

	// Stats and internal variables.
	TotalLeaves    *int `json:"-"` // For whole tree, up and down.
	MaxTotalLeaves *int `json:"-"` // For whole tree, up and down.

	// Recursive struct variable.  Last position in for easier reading.
	Children []*Tree `json:"children,omitempty"`
}

// TODO
type TreeStats struct {
	TotalLeaves    int `json:"total_leaves,omitempty"`
	MaxTotalLeaves int `json:"max_total_leaves,omitempty"`
}

// Populate needs alg and seed set before calling, and should set size if tree
// is non-empty. Does not populate trunk.
func (t *Tree) Populate() (err error) {
	//fmt.Printf("Populate tree: %+v\n", t)
	err = t.GenTreeBranches()
	if err != nil {
		return err
	}

	if len(t.BLS) > 1 { // Recursive tree.
		t.Children = make([]*Tree, t.BLS[0])
		for i := 0; i < t.BLS[0]; i++ {
			tt := new(Tree)
			tt.Alg = t.Alg
			tt.Seed = t.Branches[i]
			tt.BLS = t.BLS[1:]
			tt.MaxTotalLeaves = t.MaxTotalLeaves
			tt.TotalLeaves = t.TotalLeaves
			tt.PathCalc = t.PathCalc

			err = tt.Populate()
			if err != nil {
				return err
			}
			t.Children[i] = tt

			if t.PathCalc {
				t.Leaves = append(t.Leaves, tt.Leaves...) // Add leaves from branch to parent's leaves.
				ps := []coze.B64{t.Seed}                  // Add parent to each child's path.
				for k, v := range tt.Paths {              // TODO need to make new unique paths, not all paths.(Use the pointer)
					paths := append(ps, *v...)
					t.Paths[k] = &paths
				}
			}
		}
	}

	t.LeavesID = make([]coze.B64, 0, len(t.Leaves)) // Preallocate for performance.
	for _, v := range t.Leaves {                    // TODO use []*[]coze.B64 for memory/processing efficiency.
		id, err := Identity(t.Alg, v)
		if err != nil {
			return err
		}
		t.LeavesID = append(t.LeavesID, id)
	}

	return t.CalcPathsID()
}

// GenTreeBranches generates a slice of digests from a seed digest for the
// current level (non-recursive).  See notes on "Tree". Resulting digests are
// not cryptographically related.
//
// Skip is the starting position for the main trunk (should typically be 0).
//
//	S─(S)─► ID
//	│
//	├─(S+0)──► D1
//	│
//	├─(S+1)──► D2
//	│
//	├─(S+2)──► D3
//	...
func (t *Tree) GenTreeBranches() (err error) {
	if len(t.Seed) != t.Alg.Size() {
		return fmt.Errorf("seed length %d does not match hashing algorithm size %d. ", len(t.Seed), t.Alg.Size())
	}
	t.ID, err = Identity(t.Alg, t.Seed)
	if err != nil {
		return err
	}
	if t.TotalLeaves == nil {
		t.TotalLeaves = new(int)
		*t.TotalLeaves = t.Skip
	}
	t.Branches = make([]coze.B64, t.BLS[0]-t.Skip)
	t.Paths = make(B64MapP)
	path := &[]coze.B64{t.Seed} // All children in a branch share the same Path (a pointer).

	for i := 0; i < len(t.Branches); i++ {
		//fmt.Printf("Current Leaves: %d\n", *t.TreeTotalLeaves)
		if t.MaxTotalLeaves != nil && *t.TotalLeaves == *t.MaxTotalLeaves { // Max condition
			return nil
		}

		node, err := Branch(t.Alg, t.Seed, i+t.Skip)
		if err != nil {
			return err
		}
		t.Branches[i] = node
		if len(t.BLS) == 1 { // Increment Leaves. (Not branches)
			//fmt.Printf("GenTreeBranches Edge of tree %t %+v \n", t.PathCalc, node)
			*t.TotalLeaves++
			if t.PathCalc {
				t.Leaves = append(t.Leaves, node)
			}
		}
		if t.PathCalc {
			t.Paths[coze.SB64(node)] = path
		}
	}

	return nil
}

// CalcPathsID converts Tree.Paths from "seed"/"private" form to "id"/"public"
// form and sets Tree.IDPaths.
func (t *Tree) CalcPathsID() (err error) {
	if !t.PathCalc {
		return
	}
	t.PathsID = make(B64MapP)
	for k, v := range t.Paths {
		kid, err := Identity(t.Alg, coze.B64(k))
		if err != nil {
			return err
		}

		var pathsId []coze.B64
		for _, v1 := range *v {
			vid, err := Identity(t.Alg, v1)
			if err != nil {
				return err
			}
			pathsId = append(pathsId, vid)
		}

		t.PathsID[coze.SB64(kid)] = &pathsId
	}
	return nil
}

// Identity is the identity function for a seed.  `S─(S)─► ID`  The result of
// the identity function is the "identifier".  All nodes are seeds from a single
// tree perspective.
func Identity(alg coze.HshAlg, b coze.B64) (coze.B64, error) {
	return coze.Hash(alg, b)
}

// Branch is the branch function for a seed.  `S─(S||byte 0)──► S0`  The result
// of this branch function is the "branch seed", which is the leaf at the tip of
// the tree.
func Branch(alg coze.HshAlg, dig coze.B64, i int) (coze.B64, error) {
	return coze.Hash(alg, append(dig, IntToBytesBE(uint64(i))...))
}

// IntToBytesLE is like binary.LittleEndian.PutUint64 except it does not include
// empty padding bytes and always returns at least one byte.   (Empty padding
// bytes on right, e.g. decimal 65,536 is [0 0 1].) As currently written, it
// only supports numbers up to 2^64 (18,446,744,073,709,551,615) which is the
// max value for uint64.
func IntToBytesLE(i uint64) []byte {
	n := 8
	b := make([]byte, n)
	binary.LittleEndian.PutUint64(b, i)
	for n > 1 && b[n-1] == 0 {
		n--
	}
	return b[:n]
}

// IntToBytesBE is like binary.BigEndian.PutUint64 except it does not include
// empty padding bytes and always returns at least one byte.  (Empty padding
// bytes on left, e.g. decimal 65,536 is bytes [1 0 0].) As currently written,
// it only supports numbers up to 2^64 (18,446,744,073,709,551,615) which is the
// max value for uint64.
func IntToBytesBE(i uint64) []byte {
	n := 8
	b := make([]byte, n)
	binary.BigEndian.PutUint64(b, i)
	j := 0
	for j < n-1 && b[j] == 0 {
		j++
	}
	return b[j:]
}

// NewTreePopulated is a simple helper for making trees.
func NewTreePopulated(alg coze.HshAlg, seed coze.B64, bls []int) (*Tree, error) {
	t := Tree{
		Alg:  alg,
		Seed: seed,
		BLS:  bls,
	}
	err := t.Populate()
	return &t, err
}

func (t Tree) String() (s string) {
	b, err := coze.MarshalPretty(t)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// B64Map is for maps with B64 as keys.  See notes on SB64.  Without this,
// defining struct values as `map[SB64][]B64` will undesirably print map keys as
// []byte and not B64 since json.Marshal looks for the key's primitive type if a
// defined custom type for the top level struct type is not present.
type B64Map map[coze.SB64][]coze.B64

// MarshalJSON implements json.Marshaler.
func (t B64Map) MarshalJSON() ([]byte, error) {
	i := 0
	l := len(t)
	s := "{"
	for k, v := range t {
		i++
		vj, err := coze.Marshal(v)
		if err != nil {
			return nil, err
		}
		s += `"` + fmt.Sprint(k) + `":` + string(vj)
		if i != l {
			s += ","
		}
	}
	return []byte(s + "}"), nil
}

// B64Map is for maps with B64 as keys.  See notes on SB64.  Without this,
// defining struct values as `map[SB64][]B64` will undesirably print map keys as
// []byte and not B64 since json.Marshal looks for the key's primitive type if a
// defined custom type for the top level struct type is not present.
type B64MapP map[coze.SB64]*[]coze.B64

// MarshalJSON implements json.Marshaler.
func (t B64MapP) MarshalJSON() ([]byte, error) {
	var s = "{"
	var i = 0
	var l = len(t)
	for k, v := range t {
		i++
		vj, err := coze.Marshal(v)
		if err != nil {
			return nil, err
		}
		s += "\"" + fmt.Sprint(k) + "\":" + string(vj)
		if i != l {
			s += ","
		}
	}
	return []byte(s + "}"), nil
}
