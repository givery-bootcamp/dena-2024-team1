import { Routes, Route } from 'react-router-dom';

import { Posts } from '../features/posts';

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<Posts />} />
    </Routes>
  );
};
