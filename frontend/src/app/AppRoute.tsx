import { Routes, Route } from 'react-router-dom';


import { Posts } from '../features/posts';
import { PostDetailPage } from '../pages/PostDetailPage';

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<Posts />} />
      <Route path="/posts/:postId" element={<PostDetailPage />} />
    </Routes>
  );
};
