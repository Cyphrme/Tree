# Tree
-------------------------------------

Package tree constructs a digest tree designed for one-way deterministic data
structures based on a single seed.  The unique property of digest tree is that
although branch digests are cryptographically related to a parent seed, the
branches themselves cannot be used to calculate the value of another branch
unless the parent seed is disclosed.  Sibling branches appear cryptographically
unrelated and branches are unaware of the parent tree and peers unless
disclosed.

Tree is a recursive datastructure.  A tree may contain other children trees and
infinite children may be calculated from a single seed. A digest tree travels
from the seed to the leaves which is opposite to a Merkel tree which travels
from leaves to a root.  

This data structure is a digest tree and not a hash tree as the digest value is
utilized. If needing further specificity we've considered dubbing this
particular form a "lightning digest" tree due to the jagged stair step and
branching shape of the structure.

### Naming
A **seed** is a single digest used to generate the tree.   A **branch** is a
tree made from the digest of a parent seed.  A branch with no recursions is a
**leaf** (a branch with inner tree size of 0 is a leaf). Seeds, branches and
leaves are **nodes**.  Trees that are produced from a **parent** tree are
**children** trees.  Branches that share a common parent are siblings.  

Terms are relative since tree is recursive.  Without considering a parent,
relative to itself, a branch's digest is a seed.  While one tree may term a
digest a seed, but from the parent tree's perspective, that seed is a branch.
For a tree of size 1, the seed is a branch and a leaf.  

**Depth sizes** denotes how large each branch level is from a starting seed. See
the example below for depth sizes of `[5,100,7,30]`.  Depth sizes are ordered
from the seed to the leaves.  This library assumes symmetrical trees.  
 

### Structure
The digest calculation step appends byte(s) denoting branch number uses
little endian, where the least significant byte is to the left. Note that the
"Identity" of the seed is not apart of the branch.

	S─(S)─► ID
	│
	├─(S||byte 0)──► S0 (Branch Digest 0)
	│
	├─(S||1)──► S1
	│
	├─(S||2)──► S2
	...

`0`, `1`, and `2` represent byte 0, byte 1, and byte 2.  Multiple bytes are
encoded little endian (see test).

Recursive, using a relative:

	S─(S)─► ID
	│
	├─(S||byte 0)──► S0 ─(S0)─► S0.ID
	│                │
	│                ├─(S0||byte 0)──► S0.0
	│                │
	│                ├─(S0||byte 1)──► S0.1
	│                ...
	│
	├─(S||byte 1)──► S1 ─(S1)─► S1.ID
	│                │
	│                ├─(S1||byte 0)──► S1.0
	...              ...

### Example: Depth sizes [5,100,7,30]

Branch levels: 4

Number of nodes in each branch level:

	Level 0: 1       (The seed)
	Level 1: 5
	Level 2: 500     (5*100)
	Level 3: 3,500   (5*100*7)
	Level 4: 105,000 (5*100*7*30) (Leaves)


Total number of nodes in each branch level:

	Level 0: 105,001  (1 + 5 * 100 * 7 * 30)    (The seed)
	Level 1: 21,001   (1 + 100 * 7 * 30)
	Level 2: 211      (1 + 7 * 30)
	Level 3: 31       (1 + 30)
	Level 4: 1        (1)                       (Leaf)



----------------------------------------------------------------------
# Attribution, Trademark notice, and License
DigestTree is released under The 3-Clause BSD License. 

"Cyphr.me" is a trademark of Cypherpunk, LLC. The Cyphr.me logo is all rights
reserved Cypherpunk, LLC and may not be used without permission.