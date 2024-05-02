INSERT INTO roles (role_name, admin) 
VALUES 
    ('SysAdmin', true),
    ('Owner', false),
    ('Manager', false),
    ('Member', false);

INSERT INTO permissions (permission_name) 
VALUES 
    ('ManageSystem'),
    ('ManageGroup'),
    ('ManageGroupMembers'),
    ('ManageRoles'),
    ('CreateContent'),
    ('EditContent'),
    ('DeleteContent'),
    ('ViewContent'),
    ('CommentOnContent'),
    ('ShareContent');

INSERT INTO role_permissions (role_id, permission_id)
VALUES
    -- SysAdmin
    (1, 1), -- ManageSystem
    (1, 2), -- ManageGroup
    (1, 3), -- ManageGroupMembers
    (1, 4), -- ManageRoles
    (1, 5), -- CreateContent
    (1, 6), -- EditContent
    (1, 7), -- DeleteContent
    (1, 8), -- ViewContent
    (1, 9), -- CommentOnContent
    (1, 10), -- ShareContent

    -- Owner
    (2, 2), -- ManageGroup
    (2, 3), -- ManageGroupMembers
    (2, 4), -- ManageRoles
    (2, 5), -- CreateContent
    (2, 6), -- EditContent
    (2, 7), -- DeleteContent
    (2, 8), -- ViewContent
    (2, 9), -- CommentOnContent
    (2, 10), -- ShareContent

    -- Manager
    (3, 3), -- ManageGroupMembers
    (3, 5), -- CreateContent
    (3, 6), -- EditContent
    (3, 7), -- DeleteContent
    (3, 8), -- ViewContent
    (3, 9), -- CommentOnContent
    (3, 10), -- ShareContent

    -- Member
    (4, 5), -- CreateContent
    (4, 8), -- ViewContent
    (4, 9), -- CommentOnContent
    (4, 10); -- ShareContent