import { Routes, Route } from "react-router-dom";

import { PostListPage } from "~/pages/PostListPage";
import { PostDetailPage } from "~/pages/PostDetailPage";
import { CreatePostPage } from "~/pages/CreatePostPage";

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<PostListPage />} />
      <Route path="/posts/:postId" element={<PostDetailPage />} />
      <Route path="/create_post" element={<CreatePostPage />} />
    </Routes>
  );
};
