CREATE TYPE sex AS ENUM ('male', 'female', 'unknown');

CREATE TABLE user_filters (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
user_id UUID NOT NULL,
search_type_id UUID NOT NULL ,
sex sex NOT NULL DEFAULT 'unknown',
use_target_id UUID NOT NULL ,
age_from int NULL,
age_to int NULL,
height_from int NULL,
height_to int NULL,
created_at timestamptz NOT NULL DEFAULT now(),
updated_at timestamptz NOT NULL DEFAULT now()
  CONSTRAINT fk_targets FOREIGN KEY (search_type_id) REFERENCES search_type(id)
);

CREATE UNIQUE INDEX ux_user_filters_user_id ON user_filters(user_id);



CREATE TABLE user_filters_tags (
user_filter_id UUID NOT NULL,
tag_id UUID NOT NULL,
created_at timestamptz NOT NULL DEFAULT now(),
PRIMARY KEY (user_filter_id, tag_id)
);

CREATE INDEX idx_user_filters_tags_tag_id ON user_filters_tags(tag_id);



CREATE TABLE search_types (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
code VARCHAR(50) NOT NULL DEFAULT 'romantic'
name VARCHAR(50) NOT NULL 
);

CREATE UNIQUE INDEX ux_user_filters_user_search_type
ON user_filters(user_id, search_type_id);