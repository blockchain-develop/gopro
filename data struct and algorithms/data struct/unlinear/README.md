map 是c++ key value [pair<key,value>]存储容器，底层使用红黑树， 结构具有自动排序功能， 1key对应1value， 查找O(logN)

unordered_map 是c++ 另外一种 key value 存储容器，底层使用哈希表，不具备排序功能，1key对应1value 查找O(1)

multimap 是c++ key value 存储容器，底层使用红黑树， 结构具有自动排序功能，允许一个key对应多个value 查找O(logN)

set 是c++ key 存储容器，底层使用红黑树， 结构具有自动排序功能， 不予许key重复 查找O(logN)

unordered_set 是c++ 另外一种 key存储容器，底层使用哈希表，不具备排序功能，不予许key重复 查找O(1)

multiset 是c++ 另外一种 key存储容器，底层使用哈希表，具备排序功能，予许key重复 查找O(logN)