import { useEffect } from 'react';
import { Link } from 'react-router-dom';

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
    <Link to={`/post/${post.id}`} key={index} className="flex justify-between">
      <div className="font-semi-bold text-4xl">
        {post.title}
      </div>
      <div>
        <div className='font-semi-bold text-gray-100 text-lg'>
          {post.userName}
        </div>
        <div className='text-gray-200 text-sm'>
          更新日時: {post.updatedAt}
        </div>
      </div>
    </Link>
    ))}
  </div>
  );
}
