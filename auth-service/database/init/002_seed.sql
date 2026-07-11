INSERT INTO services (name, active)
VALUES
    ('orders-service', TRUE),
    ('profile-service', TRUE),
    ('disabled-service', FALSE)
ON CONFLICT (name) DO NOTHING;