import { useCallback , FormEvent } from "react";
import { useNavigate } from "react-router-dom";

import { Container } from "~/shared/components/Container";
import { Button } from "~/shared/components/Button";
import { postApi } from "~/shared/services/API";

export function CreatePostPage() {
  const navigate = useNavigate();

  const handleSubmit = useCallback(async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const form = event.currentTarget;
    const formData = new FormData(form);
    const title = formData.get("title") as string;
    const body = formData.get("body") as string;
    const user_id = 1;
    const response = await postApi.postPost({
      title,
      body,
      user_id,
    } , {
      withCredentials: true,
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
