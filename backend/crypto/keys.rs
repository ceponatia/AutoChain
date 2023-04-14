use ed25519_dalek::{Keypair, PublicKey, SecretKey, Signature, Signer, Verifier};
use rand::rngs::OsRng;

pub fn generate_keypair() -> Keypair {
    let mut csprng = OsRng {};
    Keypair::generate(&mut csprng)
}

pub fn extract_public_key(keypair: &Keypair) -> PublicKey {
    keypair.public
}

pub fn serialize_public_key(public_key: &PublicKey) -> Vec<u8> {
    public_key.to_bytes().to_vec()
}

pub fn deserialize_public_key(bytes: &[u8]) -> Result<PublicKey, ed25519_dalek::SignatureError> {
    PublicKey::from_bytes(bytes)
}

pub fn sign_message(keypair: &Keypair, message: &[u8]) -> Signature {
    keypair.sign(message)
}

pub fn verify_signature(
    public_key: &PublicKey,
    message: &[u8],
    signature: &Signature,
) -> Result<(), ed25519_dalek::SignatureError> {
    public_key.verify(message, signature)
}