import { useCallback , FormEvent } from "react";
import { useNavigate } from "react-router-dom";

import { Container } from "~/shared/components/Container";
import { Button } from "~/shared/components/Button";
import { PostApi } from "~/generated";

export function CreatePostPage() {
  const navigate = useNavigate();

  const handleSubmit = useCallback(async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const form = event.currentTarget;
    const formData = new FormData(form);
    const title = formData.get("title") as string;
    const body = formData.get("body") as string;
    const user_id = 1;
    const postApi = new PostApi();
    const response = await postApi.postPost({
      title,
      body,
      user_id,
    } , {
      withCredentials: true,
    });
    console.log(response);
    if (response.status === 201) {
      console.log("Created post successfully!");
      navigate("/");
    }

  }, []);
  
  return (
    <Container className="mt-20">
      <div>
        <form onSubmit={handleSubmit} >
          <div>
            <label className="text-2xl font-bold">
              タイトル
            </label>
            <br/>
            <input type="text" id="title" name="title" className="w-full rounded border border-gray-200 text-2xl" />
          </div>
          <br/>
          <div>
            <label className="text-2xl font-bold">
              本文
            </label>
            <br/>
            <textarea id="body" name="body" className="h-64 w-full rounded border border-gray-200 text-2xl" />
          </div>
          <br/>
          <Button className="w-full" type="submit">投稿する</Button>
        </form>
      </div>
    </Container>
  );
}
