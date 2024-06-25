import { useEffect } from "react";
import { useParams } from "react-router-dom";

import { PostActionButtons } from "~/features/posts/PostActionButtons";
import { PostContent } from "~/features/posts/PostContent";
import { PostHeading } from "~/features/posts/PostHeading";
import { Container } from "~/shared/components/Container";
import { useAppSelector, useAppDispatch } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { postApi } from "~/shared/services/API";
import { formatDateTime } from "~/shared/utils";

export function PostDetailPage() {
  const { post } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();
  const { postId } = useParams<{postId: string}>();

  useEffect(() => {
    dispatch(APIService.getPost(Number(postId)));
  }, [dispatch]);

  const handleDelete = async () => {
    const response = await postApi.deletePost(Number(postId) , {
      withCredentials: true,
    });
    if (response.status === 204) {
      console.log("Deleted post successfully!");
      window.location.href = "/";
    }
  };

  if (!post) return <p>Loading...</p>;
  return (
    <Container>
      <div className="flex flex-col gap-6">
        <div className="mt-10 flex flex-col gap-7">
          <PostHeading
            title={post.title}
            createdAt={formatDateTime(post.createdAt)}
            updatedAt={formatDateTime(post.updatedAt)}
          />
          <hr className="border-border" />
          <PostContent body={post.body} username={post.userName} />
        </div>
        <PostActionButtons postId={post.id} onDelete={handleDelete} />
      </div>
    </Container>
  );
}
