ALTER TABLE webinar_users RENAME webinar_id TO online_id;
ALTER TABLE webinar_templates ADD COLUMN class_id INTEGER REFERENCES classes(id);
ALTER TABLE webinars DROP class_id;
ALTER TABLE webinar_templates ADD COLUMN img TEXT;
