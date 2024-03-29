# Tree
-------------------------------------

Package tree constructs a digest tree designed for one-way deterministic data
structures based on a single seed.  A unique property of digest tree is that
although branch digests are cryptographically related to a parent seed, a
branches cannot be used to calculate the value of a sibling branch. Branches
appear cryptographically unrelated to the parent tree and siblings unless the
seed is disclosed.

Tree is a recursive datastructure.  A tree may contain other children trees and
infinite children may be calculated from a single seed. A digest tree travels
from the seed to the leaves which is opposite to a Merkle tree which travels
from leaves to a root.  

This data structure is a digest tree and not a hash tree as the digest value is
utilized. If needing further specificity we've considered dubbing this
particular form a "lightning digest" tree due to the jagged stair step and
branching shape of the structure.

### Naming
A **seed** is a single digest used to generate the tree.   A **branch** is a
tree made from the digest of a parent seed.  A branch with no recursions is a
**leaf** (a branch with child tree size of 0 is a leaf). Seeds, branches and
leaves are **nodes**.  Trees that are produced from a **parent** tree are
**children** trees.  Branches that share a common parent are siblings.  

Terms are relative since tree is recursive.  Without considering a parent,
relative to itself, a branch's digest is a seed.  While one tree may term a
digest a seed, but from the parent tree's perspective, that seed is a branch.
For a tree of size 1, the seed is a branch and a leaf.  

**Branch level sizes** (BLS) denotes the size of each branch level.  The
first integer is for the current level and subsequent integers are for each
child level. See the example below for BLS of `[5,100,7,30]`.  Branch sizes are
ordered from the seed to the leaves.  This library assumes symmetrical trees,
although non-symmetrical trees are legitimate.  
 

### Structure
Input seed is a random number of the same size as the digest algorithm's output,
termed the hash size. The identity (id) of the seed is the digest of the seed.
Branch digests are calculated by appending a sequential nonce, encoded as big
endian byte(s), to the seed and hashing.  

	S ─(S)────────► ID
	│
	├─(S||byte 0)─► S0 (Branch Digest 0)
	│
	├─(S||byte 1)─► S1 (Branch Digest 1)
	│
	├─(S||byte 2)─► S2 (Branch Digest 2)
	...
	├─(S||byte 255) ───────► S255 (Branch Digest 255)
	│
	├─(S||byte 1||byte 0) ─► S256 (Branch Digest 256)

Recursive:

	S─(S)─────────► ID
	│
	├─(S||0)───────► S0 ─(S0)─► S0.ID
	│                │
	│                ├─(S0||0)──► S0.0
	│                │
	│                ├─(S0||1)──► S0.1
	│                ...
	│
	├─(S||1)───────► S1 ─(S1)─► S1.ID
	│                │
	│                ├─(S1||0)──► S1.0
	...              ...

Because the nonce is sequential, any branch can quickly be calculated by
position in the tree.   Note that the identity of the seed is not apart of the
branch.

### Example: Branch level sizes [5,100,7,30]

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


# What is a digest tree useful for?
We're using it for deterministic ticket generation.  From a single seed, all
possible tickets are easily calculable.  Instead of storing pre-calculated
tickets in a database, only a seed needs to be stored.  As long as the path to
the ticket is given, an agent with knowledge of the private seed can later
confirm that a given ticket is legitimate.  Although this is possible with a
simple digest construction with a seed concatenated with a serial byte, we've
found the need for layers or levels of tickets, which the tree structure
provides.  

For a physical example, suppose a business sells rolls of carnival tickets to
distributors, who then sell individual tickets to customers.  From a single seed
many rolls of tickets can be generated, and each roll of tickets has many
tickets.  A ticket roll may have it's private seed disclosed, so the agent in
possession can calculate it's own tickets, or only the ticket roll seed id could
be disclosed, which does not allow the agent in possession of the ticket roll to
calculate the contained tickets.  


# Significand and Endianness
The digest calculation step appends an integer, a nonce denoting branch number,
encoded as big endian byte(s), e.g. byte 0, byte 1, etc.  The minimal number of
bytes needed to encode the nonce are the **significand bytes**.  Only the
significand bytes are valid, and non-significand bytes, named padding bytes, are
invalid. 

Big endian has the most significant byte is to the left. As is typical byte
paradigm binary counting, at the 255 and 256 bit boundary, an additional byte is
needed to represent 256, and additional bytes needed for each boundary
thereafter. For big endian, preceding zero'd bytes, that is bytes on the left,
are padding bytes.  This is the same as decimal integers where padding precedes the
significand.  

Little endian has the least significant byte is to the left. As is typical
byte paradigm binary counting, at the 255 and 256 bit boundary.  An additional
byte is needed to represent 256, and additional bytes needed for each boundary
thereafter. For little endian, succeeding zero'd bytes, that is bytes are on the
right, are invalid.  This is opposite to decimal integers where padding precedes the
significand.  



----------------------------------------------------------------------
# Attribution, Trademark notice, and License
DigestTree is released under The 3-Clause BSD License. 

"Cyphr.me" is a trademark of Cypherpunk, LLC. The Cyphr.me logo is all rights
reserved Cypherpunk, LLC and may not be used without permission.