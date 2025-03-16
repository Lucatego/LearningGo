-- Script for the sequence table

-- Redo User
-- Note: Use only when the other tables are empty, like UserConversation or UserMessage.
-- If the other tables have elements, then another script should be used.
DELETE FROM User;
DELETE FROM sqlite_sequence WHERE name = 'User';
