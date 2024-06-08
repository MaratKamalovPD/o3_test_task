INSERT INTO public.post (user_id, title, "content", are_comments_disabled)
VALUES  (1, 'First Post', 'This is my very first post.', false),
        (2, 'Another Post', 'Here is another post for demonstration.', true),
        (3, 'Post with Comments', 'Lets see what people have to say about this.', false),
        (4, 'Empty Post', ' ', false),
        (5, 'Long Content Post', 'This is a very long piece of content designed to test how much text we can fit into a single post.', false),
        (6, 'Special Characters in Title', 'Using special characters like &quot; and &lt;&gt; in titles can be fun!', false),
        (7, 'No Comment Post', 'Comments are disabled on this post. Feel free to share elsewhere.', true),
        (8, 'Short Content Post', 'A brief update.', false),
        (9, 'Numbers in Title', 'Including numbers in your post titles can make them stand out.', false),
        (10, 'HTML in Content', '<p>This is a <strong>post</strong> with <em>HTML tags</em>.</p>', false),
        (11, 'Multiple Lines Post', 'This post spans across several lines to demonstrate.\n\nLine 1\nLine 2\nLine 3', false),
        (12, 'New Line Character', 'This post demonstrates the use of the newline character (\n).', false);


INSERT INTO public.comment (post_id, parent_comment_id, user_id, content)
VALUES  (1, NULL, 1, 'Great post'),
        (1, 1, 2, 'Thanks for sharing.'),
        (1, 2, 3, 'I agree with both of you.'),
        (2, NULL, 4, 'Interesting read.'),
        (2, 3, 5, 'That was insightful.'),
        (3, 4, 6, 'Agreed, it was quite enlightening.'),
        (4, 5, 7, 'This is a great discussion going on here.'),
        (5, 8, 1, 'I couldn''t agree more.'),
        (6, NULL, 2, 'This post really made me think.'),
        (7, 10, 3, 'Excellent point, well said.'),
        (8, NULL, 4, 'This is a thoughtful comment.'),
        (9, 6, 5, 'I completely disagree with that perspective.'),
        (10, 7, 1, 'Well said, I appreciate the insight.'),
        (11, NULL, 2, 'This post has opened up a whole new perspective for me.'),
        (12, 9, 3, 'Your analysis was spot-on.'),
        (1, NULL, 1, 'This post has sparked a lot of thought.'),
        (2, 11, 2, 'Im glad to hear that.'),
        (3, 12, 3, 'Indeed, it was quite insightful.'),
        (4, NULL, 4, 'This post really resonated with me.'),
        (5, 14, 5, 'Thank you for sharing your thoughts.');


