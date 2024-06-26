import { Routes, Route, Navigate } from "react-router-dom";

import { SketchListPage } from "~/pages/SketchListPage";
import { SketchCreatePage } from "~/pages/SketchCreatePage";
import { SignInPage } from "~/pages/SignInPage";
import { SignUpPage } from "~/pages/SignUpPage";
import { SignOutPage } from "~/pages/SignOutPage";
import { PostDetailPage } from "~/pages/PostDetailPage";
import { CreatePostPage } from "~/pages/CreatePostPage";
import { PostEditPage } from "~/pages/PostEditPage";
import { PostListPage } from "~/pages/PostListPage";

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<SketchListPage />} />
      <Route path="/sketches/create" element={<SketchCreatePage />} />
      <Route path="/signout" element={<SignOutPage />} />
      <Route path="/posts" element={<PostListPage />} />
      <Route path="/posts/:postId" element={<PostDetailPage />} />
      <Route path="/posts/create" element={<CreatePostPage />} />
      <Route path="/posts/:postId/edit" element={<PostEditPage />} />
      <Route path="*" element={<Navigate to={"/"} />} />
    </Routes>
  );
};

export const AppRouteNoAuth = () => {
  return (
    <Routes>
      <Route path="/" element={<SketchListPage/>} />
      <Route path="/signin" element={<SignInPage />} />
      <Route path="/signup" element={<SignUpPage />} />
      <Route path="*" element={<Navigate to={"/"} />} />
    </Routes>
  );
};
