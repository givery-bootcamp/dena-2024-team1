import { Routes, Route } from "react-router-dom";


import { Posts } from '../pages/Posts';
import { PostDetailPage } from '../pages/PostDetailPage';


export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<Posts />} />
      <Route path="/posts/:postId" element={<PostDetailPage />} />
    </Routes>
  );
};
