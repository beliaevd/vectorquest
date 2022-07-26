ALTER TABLE webinar_users RENAME online_id TO webinar_id;
ALTER TABLE webinar_templates DROP COLUMN class_id;
ALTER TABLE webinars ADD COLUMN class_id  INTEGER REFERENCES classes(id);
ALTER TABLE webinar_templates DROP COLUMN img;
