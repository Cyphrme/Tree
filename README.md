# Tree
-------------------------------------

Tree is a digest tree designed for one-way deterministic data structures based
on a single seed.  Resulting branch digests are not cryptographically related as
one branch cannot be used to calculate the value of another branch. Tree is a
recursive datastructure and a tree may contain other trees.  A digest tree
travel direction is opposite to a Merkel tree; instead of from the leaves up to
a root, a digest tree travels from a seed to the leaves.

This data structure is a "digest tree" and not a "hash tree" because the digest
value is utilized. If needing further specificity, we've considered dubbing this
particular form a "lightning digest" tree due to the jagged stair step and
branching shape of the structure.

### Naming
The **seed** is a single digest used to generate the tree, "the trunk".   A
**branch** is a tree that's made from the digest of a parent seed.  A branch
with no recursions is a **leaf** (a branch with inner tree size 0 is a leaf).
All "seeds", "branches" and "leaves" are **nodes**.

Since tree is recursive, many of the terms are relative.  Where one tree may
term a digest "seed", from the perspective of a different tree, the "seed" may
also be a "leaf".  

**Depth sizes** denote how large each branch level is from a starting seed.  See
the example below for depth sizes of `[5,100,7,30]`.  Depth sizes is ordered
from the trunk to the leaves.  
 

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
	...             ...

### Example: Depth sizes [5,100,7,30]

Branch levels: 4


Number of nodes in each branch level:

	Level 0: 1       (The seed)
	Level 1: 5
	Level 2: 500     (5*100)
	Level 3: 3,500   (5*100*7)
	Level 4: 105,000 (5*100*7*30) (Leaves)


Total number of nodes in each tree level:

	Level 0: 1       (The seed)
	Level 1: 6       (1 + 5)
	Level 2: 506     (1 + 5 + 500)
	Level 3: 4,006   (1 + 5 + 500 + 3,500)
	Level 4: 109,006  (1 + 5 + 500 + 3,500 + 105,000) 



----------------------------------------------------------------------
# Attribution, Trademark notice, and License
DigestTree is released under The 3-Clause BSD License. 

"Cyphr.me" is a trademark of Cypherpunk, LLC. The Cyphr.me logo is all rights
reserved Cypherpunk, LLC and may not be used without permission.