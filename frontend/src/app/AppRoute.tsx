import { Routes, Route, Navigate } from "react-router-dom";

import { PostListPage } from "~/pages/PostListPage";
import { PostDetailPage } from "~/pages/PostDetailPage";
import { CreatePostPage } from "~/pages/CreatePostPage";
import { SignInPage } from "~/pages/SignInPage";
import { SignUpPage } from "~/pages/SignUpPage";
import { SignOutPage } from "~/pages/SignOutPage";

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<PostListPage />} />
      <Route path="/posts/:postId" element={<PostDetailPage />} />
      <Route path="/posts/create" element={<CreatePostPage />} />
      <Route path="/signout" element={<SignOutPage />} />
      <Route path="*" element={<Navigate to={"/"} />} />
    </Routes>
  );
};

export const AppRouteNoAuth = () => {
  return (
    <Routes>
      <Route path="/signin" element={<SignInPage />} />
      <Route path="/signup" element={<SignUpPage />} />
      <Route path="*" element={<Navigate to={"/signin"} />} />
    </Routes>
  );
};
