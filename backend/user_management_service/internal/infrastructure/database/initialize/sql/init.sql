INSERT INTO roles (role_name) VALUES
  ('SysAdmin'),
  ('Owner'),
  ('Admin'),
  ('Member');

INSERT INTO permissions (permission_name) VALUES
  ('ManageSystem'),
  ('CreateGroup'),
  ('DeleteGroup'),
  ('UpdateGroupSettings'),
  ('InviteMembers'),
  ('RemoveMembers'),
  ('AssignRoles'),
  ('CreateContent'),
  ('EditContent'),
  ('DeleteContent'),
  ('ViewContent'),
  ('CommentOnContent'),
  ('ShareContent');

INSERT INTO role_permissions (role_id, permission_id)
VALUES
  (1, 1), (1, 2), (1, 3), (1, 4), (1, 5), (1, 6), (1, 7), (1, 8), (1, 9), (1, 10), (1, 11), (1, 12), (1, 13),
  (2, 2), (2, 3), (2, 4), (2, 5), (2, 6), (2, 7), (2, 8), (2, 9), (2, 10), (2, 11), (2, 12), (2, 13),
  (3, 4), (3, 5), (3, 6), (3, 8), (3, 9), (3, 10), (3, 11), (3, 12), (3, 13),
  (4, 8), (4, 11), (4, 12), (4, 13);