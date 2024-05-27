import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from '../../shared/hooks';
import { APIService } from '../../shared/services';

export function Posts() {
  const { posts } = useAppSelector((state) => state.posts);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPosts());
  }, [dispatch]);

  return (
  <div>
    {posts?.map((post,index)=>(
    <div key={index}>
      {post.title}
    </div>
    ))}
  </div>
  );
}
