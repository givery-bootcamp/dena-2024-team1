/* tslint:disable */
/* eslint-disable */
/**
 * Simple User API
 * A simple API to retrieve user information.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface PostsGet200ResponseInner
 */
export interface PostsGet200ResponseInner {
    /**
     * 
     * @type {number}
     * @memberof PostsGet200ResponseInner
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof PostsGet200ResponseInner
     */
    title?: string;
    /**
     * 
     * @type {string}
     * @memberof PostsGet200ResponseInner
     */
    body?: string;
    /**
     * 
     * @type {string}
     * @memberof PostsGet200ResponseInner
     */
    userName?: string;
    /**
     * 
     * @type {Date}
     * @memberof PostsGet200ResponseInner
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof PostsGet200ResponseInner
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the PostsGet200ResponseInner interface.
 */
export function instanceOfPostsGet200ResponseInner(value: object): value is PostsGet200ResponseInner {
    return true;
}

export function PostsGet200ResponseInnerFromJSON(json: any): PostsGet200ResponseInner {
    return PostsGet200ResponseInnerFromJSONTyped(json, false);
}

export function PostsGet200ResponseInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostsGet200ResponseInner {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'title': json['title'] == null ? undefined : json['title'],
        'body': json['body'] == null ? undefined : json['body'],
        'userName': json['userName'] == null ? undefined : json['userName'],
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'updatedAt': json['updatedAt'] == null ? undefined : (new Date(json['updatedAt'])),
    };
}

export function PostsGet200ResponseInnerToJSON(value?: PostsGet200ResponseInner | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'title': value['title'],
        'body': value['body'],
        'userName': value['userName'],
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'updatedAt': value['updatedAt'] == null ? undefined : ((value['updatedAt']).toISOString()),
    };
}
