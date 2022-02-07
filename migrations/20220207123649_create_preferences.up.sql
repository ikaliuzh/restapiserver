create table preferences (
    user_id int not null primary key,
    tracked_tokens varchar not null,
    fiat_currency varchar not null,

    foreign key (user_id) references users(id)
);