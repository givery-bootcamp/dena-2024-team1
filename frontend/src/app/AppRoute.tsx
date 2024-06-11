import { Routes, Route } from "react-router-dom";

import { PostListPage } from "~/pages/PostListPage";
import { PostDetailPage } from "~/pages/PostDetailPage";

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<PostListPage />} />
      <Route path="/posts/:postId" element={<PostDetailPage />} />
    </Routes>
  );
};
