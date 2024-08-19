SELECT setval('object_types_object_type_id_seq', COALESCE((SELECT MAX(object_type_id)+1 FROM object_types), 1), false);
SELECT setval('columns_column_id_seq', COALESCE((SELECT MAX(column_id)+1 FROM columns), 1), false);
SELECT setval('qualities_quality_id_seq', COALESCE((SELECT MAX(quality_id)+1 FROM qualities), 1), false);
SELECT setval('roles_role_id_seq', COALESCE((SELECT MAX(role_id)+1 FROM roles), 1), false);
SELECT setval('permissions_permission_id_seq', COALESCE((SELECT MAX(permission_id)+1 FROM permissions), 1), false);
SELECT setval('users_user_id_seq', COALESCE((SELECT MAX(user_id)+1 FROM users), 1), false);