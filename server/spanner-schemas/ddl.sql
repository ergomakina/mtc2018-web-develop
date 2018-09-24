CREATE TABLE Likes (
  Uuid STRING(64) NOT NULL,
  SessionId INT64 NOT NULL,
  UserUUID STRING(64) NOT NULL,
  CreatedAt TIMESTAMP NOT NULL,
) PRIMARY KEY(Uuid);

CREATE TABLE LikesSummaries (
  Second INT64 NOT NULL,
  SessionId INT64 NOT NULL,
  ServerId STRING(64) NOT NULL,
  Likes INT64 NOT NULL,
) PRIMARY KEY(Second, SessionId, ServerId);
