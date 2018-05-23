table! {
    projects (id) {
        id -> Int4,
        title -> Nullable<Varchar>,
    }
}

table! {
    users (id) {
        id -> Int4,
        name -> Varchar,
        email -> Varchar,
        password -> Varchar,
    }
}

allow_tables_to_appear_in_same_query!(
    projects,
    users,
);
