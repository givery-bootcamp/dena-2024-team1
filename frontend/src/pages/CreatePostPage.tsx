import { useCallback , FormEvent } from "react";
import { useNavigate } from "react-router-dom";

import { Container } from "~/shared/components/Container";
import { Button } from "~/shared/components/Button";
import { postApi } from "~/shared/services/API";
import { useAppSelector } from "~/shared/hooks";

export function CreatePostPage() {
  const navigate = useNavigate();
  const { sessionUser } = useAppSelector((state) => state.sessionUser);

  const handleSubmit = useCallback(async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const form = event.currentTarget;
    const formData = new FormData(form);
    const title = formData.get("title") as string;
    const body = formData.get("body") as string;

    if (!sessionUser || !sessionUser.id) {
      console.error("User is not logged in.");
      navigate("/signin");
      return;
    }

    const user_id = sessionUser.id;
    const response = await postApi.postPost({
      title,
      body,
      user_id,
    });
    if (response.status === 201) {
      console.log("Created post successfully!");
      navigate("/");
    }

  }, []);
  
  return (
    <Container className="pt-10">
      <form onSubmit={handleSubmit} className="flex flex-col gap-8">
        <label className="flex flex-col gap-1">
          <span className="font-bold">タイトル</span>
          <input className="h-10 rounded border border-gray-200 p-2" type="text" name="title" />
        </label>
        <label className="flex flex-col gap-1">
          <span className="font-bold">本文</span>
          <textarea className="rounded border border-gray-200 p-2" name="body" />
        </label>
        <Button className="h-12 w-full" type="submit">投稿する</Button>
      </form>
    </Container>
  );
}
