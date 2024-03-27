# Redis

## 1. Record

## 2. Hash

## 3. Set

## 4. Sorted Set
- Sorted set là sự kết hợp của 2 datastructure khác là set và hash.
- Với cấu trúc là member (string) và score (int), được sắp xếp dựa trên thứ tự tăng dần của score.
## ZADD: Thêm cặp member-score vào sorted set.
```
  ZADD <id> <score> <member>
```
## ZCARD: Trả về số member có trong sorted set.
```
  ZCARD <id>
```
## ZCOUNT: Trả về số member CÓ ĐIỀU KIỆN trong sorted set.
```
  ZCOUNT <id> [(]<lower_bound> [(]<upper_bound>
```
* Với [lower_bound; upper_bound], nếu không có ý định lấy cả khoảng thì có thể thêm '('.
* Có thể option thêm -inf và inf để định nghĩa khoảng.
## ZPOPMIN, ZPOPMAX: Pop phần tử bé/lớn nhất có trong sorted set.
```
  ZPOPMIN <id> [(OPTION int x) bỏ x phần tử bé nhất]
```
## ZINCRBY: Thay đổi score
```
  ZINCRBY <id> <score> <member>
```
## ZRANGE: Lấy ra phần tử theo index.
```
  ZRANGE <id> <lower_index> <upper_index>
```
* Có thể lọc ra phần tử theo score
```
  ZRANGE <id> -inf +inf BYSCORE [WITHSCORES (include thêm score)] [(REV) reversed sort] [(LIMIT) (offset) (numOfElemnt)]
```
