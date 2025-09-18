# xk6-crypto-scrypt

A public `xk6` extension for scrypt hashing in k6 JavaScript tests.

## Installation

Build k6 with this extension:

```sh
xk6 build --with github.com/exterminatorz/xk6-crypto-scrypt
```

## Usage Example (k6 script)

```js
import scrypt from 'k6/x/secp256k1'; // If your module is registered as k6/x/secp256k1

export default function () {
  const result = scrypt.generateHash('myPassword', {
    N: 16384,
    r: 8,
    p: 1,
    keyLen: 64,
    saltLen: 16,
  });
  console.log('Scrypt Hash:', result.hash);
  console.log('Scrypt Salt:', result.salt);
  console.log('Scrypt Input:', result.input);
}
```

> - `PrivateKey` (Uint8Array/ArrayBuffer) — present for private keys
> - `PublicKey` (Uint8Array/ArrayBuffer) — present for public keys
> - `Algorithm` — `"secp256k1"`
> - `Extractable` — always `true`
> - `Type` — `"private"` | `"public"`

## Supported Methods

| Function       | Input (JS)      | Output (JS)           | Description                          |
| -------------- | --------------- | --------------------- | ------------------------------------ |
| `generateHash` | string, options | `{hash, salt, input}` | Scrypt hash with configurable params |

### Options for `generateHash`

- `N` (int): CPU/memory cost parameter (default: 16384)
- `r` (int): Block size (default: 8)
- `p` (int): Parallelization (default: 1)
- `keyLen` (int): Output key length (default: 64)
- `saltLen` (int): Salt length in bytes (default: 16)

## Running the example

```sh
# Build k6 with this extension
xk6 build --with github.com/exterminatorz/xk6-crypto-scrypt

# Run the example script
k6 run examples/k6_example.js
```

## License

MIT
