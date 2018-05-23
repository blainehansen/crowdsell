# diesel migration redo
export DATABASE_URL="postgres://user:asdf@localhost/crowdsell"
diesel print-schema > server/schema.rs
