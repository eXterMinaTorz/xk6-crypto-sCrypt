package scrypt

// This package exists so the module root (github.com/exterminatorz/xk6-crypto-sCrypt)
// can be imported directly by xk6's --with flag. It blank-imports the real
// implementation in the `sCrypt` subpackage which performs the actual
// module registration in its init().

import _ "github.com/exterminatorz/xk6-crypto-scrypt/scrypt"
