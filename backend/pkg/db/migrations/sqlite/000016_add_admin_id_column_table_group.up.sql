ALTER TABLE groups
ADD admin_id UUID REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE;