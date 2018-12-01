
class MerkleSumTree {
    constructor (leaves) {
        this.checkConsecutive(leaves);

        this.buckets = leaves.map(leaf => leaf.getBucket());
        let buckets = this.buckets.slice();
        while (buckets.length !== 1) {
            const newBuckets = [];
            while (buckets.length > 0) {
                if (buckets.length >= 2) {
                    const b1 = buckets.shift();
                    const b2 = buckets.shift();
                    const size = b1.size + b2.size;
                    const hashed = utils.soliditySha3(
                        { type: 'uint64', value: b1.size },
                        { type: 'bytes32', value: b1.hashed },
                        { type: 'uint64', value: b2.size },
                        { type: 'bytes32', value: b2.hashed });
                    const b = new Bucket(size, hashed);
                    b1.parent = b2.parent = b;
                    b1.right = b2;
                    b2.left = b1;
                    newBuckets.push(b);
                } else {
                    newBuckets.push(buckets.pop());
                }
            }
            buckets = newBuckets;
        }
        this.root = buckets[0];
    }

    checkConsecutive (leaves) {
        let curr = 0;
        leaves.forEach(leaf => {
            if (leaf.rng[0] !== curr) throw new Error('Leaf ranges are invalid!');
            curr = leaf.rng[1];
        });
    }

    getRoot () {
        return this.root;
    }

    getProof (index) {
        var curr = this.buckets[index];
        const proof = [];
        while (curr.parent) {
            const right = curr.right;
            const bucket = curr.right ? curr.right : curr.left;
            curr = curr.parent;
            proof.push(new ProofStep(bucket, right));
        }
        return proof;
    }

    verifyProof (root, leaf, proof) {

        // Validates the supplied `proof` for a specific `leaf` according to the
        // `root` bucket of Merkle-Sum-Tree.
        const leftProofStepArr = proof.map(s => !s.right ? s.bucket.size : 0);
        const rightProofStepArr = proof.map(s => s.right ? s.bucket.size : 0);
        const rng = [sum(leftProofStepArr), (root.size - sum(rightProofStepArr))];

        // Supplied steps are not routing us to the range specified.
        // TODO: this needs to be an arr comparison, right now the range arrays are never equal
        if (rng[0] !== leaf.rng[0] || rng[1] !== leaf.rng[1]) return false;

        let curr = leaf.getBucket();

        let hashed;
        proof.forEach(step => {
            if (step.right) {
                hashed = utils.soliditySha3(
                    { type: 'uint64', value: curr.size },
                    { type: 'bytes32', value: curr.hashed },
                    { type: 'uint64', value: step.bucket.size },
                    { type: 'bytes32', value: step.bucket.hashed });
            } else {
                hashed = utils.soliditySha3(
                    { type: 'uint64', value: step.bucket.size },
                    { type: 'bytes32', value: step.bucket.hashed },
                    { type: 'uint64', value: curr.size },
                    { type: 'bytes32', value: curr.hashed });
            }
            curr = new Bucket(curr.size + step.bucket.size, hashed);
        });

        return curr.size === root.size && curr.hashed === root.hashed;
    }
}