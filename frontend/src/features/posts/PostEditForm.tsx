import { useCallback , FormEvent , useEffect , useState } from "react";
import { useNavigate , useParams } from "react-router-dom";

import { useAppSelector, useAppDispatch } from "~/shared/hooks";
import { Button } from "~/shared/components/Button";
import { postApi } from "~/shared/services/API";
import { APIService } from "~/shared/services";

export const PostEditForm = () => {
  const { post } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const { postId } = useParams<{postId: string}>();
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  console.log(title);
  console.log(body);
  const handleTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value);
  };

  const handleBodyChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setBody(event.target.value);
  };

  if (!post) return <p>Loading...</p>;
  useEffect(() => {
    dispatch(APIService.getPost(Number(postId)));
    //setTitle(post.title);
    //setBody(post.body); 
  }, []);

  const handleSubmit = useCallback(async (event: FormEvent<HTMLFormElement>) => {
    //console.log(title);
    //console.log(body);
    //console.log(postId);
    event.preventDefault();
    //const form = event.currentTarget;
    //const formData = new FormData(form);
    //const title = formData.get("title") as string;
    //const body = formData.get("body") as string;
    const response = await postApi.putPost(
      parseInt(postId ?? ""),
      {
        title: "固定タイトル",
        body: "固定本文",
      },
      { 
        withCredentials: true,
      },
    );
    //console.log(title);
    //console.log(body);
    //console.log(response.status);
    if (response.status === 201) {
      navigate("/");
    }

  }, []);


  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-8">
      <label className="flex flex-col gap-1">
        <span className="font-bold">タイトル</span>
        <input className="h-10 rounded border border-gray-200 p-2" type="text" name="title" onChange={handleTitleChange} value={title}/>
      </label>
      <label className="flex flex-col gap-1">
        <span className="font-bold">本文</span>
        <textarea className="rounded border border-gray-200 p-2" name="body" onChange={handleBodyChange} value={body}/>
      </label>

      <Button className="h-12" type="submit">更新する</Button>
    </form>
  );
};
