import scrypt from 'k6/x/scrypt';

export default function () {
  // --- Scrypt Hash Example ---
  const scryptResult = scrypt.generateHash('myPassword', {
    N: 16384,
    r: 8,
    p: 1,
    keyLen: 64,
    saltLen: 16,
  });
  console.log('Scrypt Hash:', scryptResult.hash);
  console.log('Scrypt Salt:', scryptResult.salt);
  console.log('Scrypt Input:', scryptResult.input);
}
import scrypt from 'k6/x/scrypt';

export default function () {
  // --- Scrypt Hash Example ---
  const scryptResult = scrypt.generateHash('myPassword', {
    N: 16384,
    r: 8,
    p: 1,
    keyLen: 64,
    saltLen: 16,
  });
  console.log('Scrypt Hash:', scryptResult.hash);
  console.log('Scrypt Salt:', scryptResult.salt);
  console.log('Scrypt Input:', scryptResult.input);
}
