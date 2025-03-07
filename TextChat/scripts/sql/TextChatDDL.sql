-- The DDL tables used to create the database.

CREATE TABLE "User" ( 
	"id" INTEGER NOT NULL UNIQUE, 
	"username" TEXT NOT NULL UNIQUE, 
	"password" TEXT NOT NULL, 
	PRIMARY KEY("id" AUTOINCREMENT) 
);

CREATE TABLE "Conversation" ( 
	"id" INTEGER NOT NULL UNIQUE, 
	"title" INTEGER NOT NULL, 
	"description" INTEGER, 
	PRIMARY KEY("id" AUTOINCREMENT) 
);

CREATE TABLE "Message" ( 
	"id" INTEGER NOT NULL, 
	"conversation_id" INTEGER NOT NULL, 
	"message" TEXT NOT NULL, 
	"time_send" TEXT NOT NULL DEFAULT (datetime('now', 'localtime')), 
	PRIMARY KEY("id","conversation_id"), 
	FOREIGN KEY("conversation_id") REFERENCES "Conversation"("id") 
);

CREATE TABLE "UserMessage" ( 
	"user_id" INTEGER NOT NULL, 
	"message_id" INTEGER NOT NULL, 
	"convesation_id" INTEGER NOT NULL, 
	"is_sender" INTEGER NOT NULL DEFAULT 0, 
	"read" INTEGER NOT NULL DEFAULT 0, 
	PRIMARY KEY("user_id","message_id","convesation_id"), 
	FOREIGN KEY("message_id") REFERENCES "Message"("id"), 
	FOREIGN KEY("user_id") REFERENCES "User"("id") 
);

CREATE TABLE "UserConversation" ( 
	"user_id" INTEGER NOT NULL, 
	"conversation_id" INTEGER NOT NULL, 
	PRIMARY KEY("user_id","conversation_id"), 
	FOREIGN KEY("conversation_id") REFERENCES "Conversation"("id"), 
	FOREIGN KEY("user_id") REFERENCES "User"("id") 
);
