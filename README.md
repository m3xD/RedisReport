# Redis

## 1. Record

## 2. Hash

## 3. Set

## 4. Sorted Set
- Sorted set là sự kết hợp của 2 datastructure khác là set và hash.
- Với cấu trúc là member (string) và score (int), được sắp xếp dựa trên thứ tự tăng dần của score.
### ZADD: Thêm cặp member-score vào sorted set.
```
  ZADD <id> <score> <member>
```
### ZCARD: Trả về số member có trong sorted set.
```
  ZCARD <id>
```
### ZCOUNT: Trả về số member CÓ ĐIỀU KIỆN trong sorted set.
```
  ZCOUNT <id> [(]<lower_bound> [(]<upper_bound>
```
* Với [lower_bound; upper_bound], nếu không có ý định lấy cả khoảng thì có thể thêm '('.
* Có thể option thêm -inf và inf để định nghĩa khoảng.
### ZPOPMIN, ZPOPMAX: Pop phần tử bé/lớn nhất có trong sorted set.
```
  ZPOPMIN <id> [(OPTION int x) bỏ x phần tử bé nhất]
```
### ZINCRBY: Thay đổi score
```
  ZINCRBY <id> <score> <member>
```
### ZRANGE: Lấy ra phần tử theo index.
```
  ZRANGE <id> <lower_index> <upper_index>
```
* Có thể lọc ra phần tử theo score
```
  ZRANGE <id> -inf +inf BYSCORE [WITHSCORES (include thêm score)] [(REV) reversed sort] [(LIMIT) (offset) (numOfElemnt)]
```
## 5. Hyperloglog

## 6. List
### LPUSH: Thêm 1 phần tử vào đầu của list.
```
  LPUSH <id> <giá trị>
```
### RPUSH: Thêm 1 phần tử vào cuối của list.
```
  RPUSH <id> <giá trị>
```
### LLEN: Trả về độ dài của list.
```
  LLEN <id>
```
### LINDEX: Trả về giá trị của phần tử theo index.
```
  LINDEX <id> <index>
```
### LRANGE: Trả về giá trị của các phần tử theo một khoảng cho trước.
```
  LRANGE <id> <lower_bound> <upper_bound>
```
> [!WARNING]
> Lấy cả upper_bound
### LPOS: Trả về index của một phần tử.
```
  LPOS <id> <GTR>
```
- Một số tham số không bắt buộc như:
  * RANK <x> : Tìm phần tử thứ x của giá trị cho trước, trả về null nếu không tồn tại.
  * COUNT <x>: Trả về x index của giá trị tính từ đầu list.
  * MAXLEN <x>: Chỉ tìm kiếm x phần tử
### LPOP: Xoá phần tử đầu tiên.
```
  LPOP <index> [x]
```
- Tham số không bắt buộc : xoá x phần tử đầu tiên
### LPOP: Xoá phần tử cuối cùng.
```
  LPOP <index> [x]
```
- Tham số không bắt buộc : xoá x phần tử cuối cùng
### LSET: Thay đổi giá trị của một phần tử.
```
  LSET <id> <index> <giá_trị_cần_thay>
```
### LTRIM: Giữ lại các phần tử có trong khoảng cho trước, xoá các phần tử còn lại.
```
  LTRIM <id> <lower_bound> <upper_bound>
```
> [!WARNING]
> Lấy cả upper_bound
### LINSERT: Thêm một phần tử vào một vị trí bất kì trong list.
```
  LINSERT <id> <BEFORE/AFTER> <giá trị có trong list> <giá trị cần thêm>
```
### LREM: Xoá một số phần tử có trong list.
```
  LREM <id> <(+/-)số lượng phần tử> <giá trị phần tử cần xoá>
```
> [!NOTE]
> Với lựa chọn (+) trước số lượng phần tử, ta tính từ đầu list và ngược lại với (-)
