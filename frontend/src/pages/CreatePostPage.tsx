import { useCallback , FormEvent } from "react";

import { Container } from "~/shared/components/Container";
import { Button } from "~/shared/components/Button";
import { PostApi } from "~/generated";

export function CreatePostPage() {
  const handleSubmit = useCallback(async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const form = event.currentTarget;
    const formData = new FormData(form);
    const title = formData.get("title") as string;
    const content = formData.get("content") as string;

    const postApi = new PostApi();
    const responnse = await postApi.postPost({
      title,
      body: content,
      user_id: 1,
    } , {
      withCredentials: true,
    });

    if (responnse.status === 200) {
      console.log("Success");

      // 画面を再読み込みする
      window.location.href = "/";
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
            <textarea id="content" name="content" className="h-64 w-full rounded border border-gray-200 text-2xl" />
          </div>
          <br/>
          <Button className="w-full" type="submit">投稿する</Button>
        </form>
      </div>
    </Container>
  );
}
