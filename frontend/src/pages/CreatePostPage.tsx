import { useState } from "react";

import { Container } from "~/shared/components/Container";
import { APIService } from "~/shared/services";
import { Button } from "~/shared/components/Button";
import { CreatePostRequest } from "~/generated";

export function CreatePostPage() {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  const handleTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value);
  };

  const handleContentChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setContent(event.target.value);
  };
  
  return (
    <Container className="mt-20">
      <div>
        <form onSubmit={(event) => {
          console.log("createPost");
          console.log(title);
          console.log(content);
          APIService.createPost({ title: title, body: content ,user_id: 1 } as CreatePostRequest);
          event.preventDefault();
        }
        } >
          <div>
            <label className="text-2xl font-bold">
              タイトル
            </label>
            <br/>
            <input type="text" value={title} onChange={handleTitleChange} className="w-full rounded border border-gray-200 text-2xl" />
          </div>
          <br/>
          <div>
            <label className="text-2xl font-bold">
              本文
            </label>
            <br/>
            <textarea value={content} onChange={handleContentChange} className="h-64 w-full rounded border border-gray-200 text-2xl" />
          </div>
          <br/>
          <Button className="w-full" type="submit">投稿する</Button>
        </form>
      </div>
    </Container>
  );
}
