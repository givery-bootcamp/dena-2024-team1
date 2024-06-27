import { Routes, Route, Navigate } from "react-router-dom";

import { SketchListPage } from "~/pages/SketchListPage";
import { SketchCreatePage } from "~/pages/SketchCreatePage";
import { SignInPage } from "~/pages/SignInPage";
import { SignUpPage } from "~/pages/SignUpPage";
import { SignOutPage } from "~/pages/SignOutPage";

export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<SketchListPage />} />
      <Route path="/sketches/create" element={<SketchCreatePage />} />
      <Route path="/signout" element={<SignOutPage />} />
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
