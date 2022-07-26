CREATE TABLE admin_category_links(
     id SERIAL PRIMARY KEY ,
     user_id INTEGER REFERENCES admin_users(id),
     category_id INTEGER REFERENCES admin_user_categories(id)
)
