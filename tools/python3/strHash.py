import os, sys, hashlib

# algs = hashlib.algorithms_available
algs = "blake2b blake2s md5 sha1 sha224 sha256 sha3_224 sha3_256 sha3_384 sha3_512 sha384 sha512"
algs = algs.split()
msg = "Please provide algorithm name and string.\n  algorithms: %s" % " ".join(algs)

if len(os.sys.argv) != 3:
    print(msg, file=sys.stderr)
    os.sys.exit(2)

alg, string = os.sys.argv[1].lower(), os.sys.argv[2]

if alg not in algs:
    print("invalid hash algorithms", file=sys.stderr)
    os.sys.exit(1)

m = hashlib.__getattribute__(alg)()
m.update(string.encode(encoding='utf-8'))
strHash = m.hexdigest()

print(strHash+"\n"+strHash.upper())
