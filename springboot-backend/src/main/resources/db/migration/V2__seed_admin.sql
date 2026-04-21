INSERT INTO admin_users(id, name, email, password_hash, role)
VALUES (
    gen_random_uuid(),
    'Super Admin',
    'admin@fds.local',
    '$2a$10$BRR83VsDvFvYMyaGx6h8/eK8Li23hOoQD/Cgr2QrXQt4QOg9QWnGG',
    'super_admin'
)
ON CONFLICT (email) DO NOTHING;
