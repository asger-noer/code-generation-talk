-- name: GetAllReminders :many
SELECT * FROM reminders
ORDER BY created ASC;

-- name: GetRemindersWithId :many
SELECT * FROM reminders 
WHERE id = ANY(@ids::uuid[])
ORDER BY created ASC;
