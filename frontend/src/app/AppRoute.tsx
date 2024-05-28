import { Routes, Route } from "react-router-dom";

import { HelloWorld } from "~/features/helloworld";
import { PostDetailPage } from "~/pages/PostDetailPage";

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<HelloWorld />} />
      <Route path="/posts/:postId" element={<PostDetailPage />} />
    </Routes>
  );
};
