CREATE TABLE IF NOT EXISTS GameSession (
    SessionId     uuid primary key not null,
	UserName      varchar(40),
	StartDate     timestamp,
	EndDate       timestamp,
	MaxToTake     integer,
	InitialAmount integer,
	MatchesLeft   integer,
	TurnHistory   integer[]
);