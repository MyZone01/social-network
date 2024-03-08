ALTER TABLE notifications
 ADD CONSTRAINT  addMoreType CHECK(type = 'follow_test');