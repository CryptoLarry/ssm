CREATE TABLE IF NOT EXISTS ssm (
  txid TEXT NOT NULL,
  address TEXT NOT NULL,
  encrypted_payload TEXT,
  encryption_version TEXT,
  primary key (txid, address)
);