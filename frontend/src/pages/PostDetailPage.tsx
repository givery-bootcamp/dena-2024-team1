import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import { PostActionButtons } from "~/features/posts/PostActionButtons";
import { PostContent } from "~/features/posts/PostContent";
import { PostHeading } from "~/features/posts/PostHeading";
import { Button } from "~/shared/components/Button";
import { Container } from "~/shared/components/Container";
import { Modal, ModalBody, ModalFooter, ModalHeader } from "~/shared/components/Modal";
import { useAppSelector, useAppDispatch } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { postApi } from "~/shared/services/API";
import { formatDateTime } from "~/shared/utils";

export function PostDetailPage() {
  const { post } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();
  const { postId } = useParams<{postId: string}>();
  const [isOpen, setIsOpen] = useState(false);

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

  const handleModalOpen = () => {
    setIsOpen(true);
  };

  const handleModalClose = () => {
    setIsOpen(false);
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
        <PostActionButtons postId={post.id} onDelete={handleModalOpen} />
      </div>
      <Modal isOpen={isOpen} onClose={handleModalClose}>
        <ModalHeader>本当に削除しますか？</ModalHeader>
        <ModalBody>
          削除した投稿は復元できません
        </ModalBody>
        <ModalFooter>
          <Button
            variant="alert"
            onClick={handleDelete}
          >
            削除する
          </Button>
          <Button
            variant="secondary"
            onClick={handleModalClose}
          >
            キャンセル
          </Button>
        </ModalFooter>
      </Modal>
    </Container>
  );
}
