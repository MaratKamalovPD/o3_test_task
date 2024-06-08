CREATE INDEX IF NOT EXISTS idx_post_id ON public.comment(post_id);
CREATE INDEX IF NOT EXISTS idx_parent_comment_id ON public.comment(parent_comment_id);
