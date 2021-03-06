/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cache

import (
	"net/http"
	"time"

	patrickmnGoCache "github.com/patrickmn/go-cache"
)

type GoCache struct {
	ttl         time.Duration
	tti         time.Duration
	rootLibrary *patrickmnGoCache.Cache
}

func NewGoCache(ttl int32, tti int32) GoCache {
	c := patrickmnGoCache.New(time.Duration(ttl)*time.Second, time.Duration(tti)*time.Second)

	gc := GoCache{
		ttl:         time.Duration(ttl) * time.Second,
		tti:         time.Duration(tti) * time.Second,
		rootLibrary: c,
	}

	return gc
}

func (c GoCache) Get(key string) *http.Response {
	item, found := c.rootLibrary.Get(key)
	if found {
		itemCopy := CopyResponse(item.(*http.Response))
		return itemCopy
	}

	return nil
}

func (c GoCache) Set(key string, value *http.Response) {
	c.rootLibrary.Set(key, value, c.ttl)
}

func (c GoCache) Delete(key string) {
	c.rootLibrary.Delete(key)
}

func (c GoCache) Clear() {
	c.rootLibrary.Flush()
}

func (c GoCache) Has(key string) bool {
	_, found := c.rootLibrary.Get(key)
	return found
}
