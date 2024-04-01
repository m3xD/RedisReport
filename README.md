# Redis

## 1. Record

## 2. Hash
- Lưu trữ dữ liệu dưới dạng (key - value)
### HSET: thêm hoặc cập nhật dữ liệu có trong hash.
```
  HSET <id> [(Các cặp key-value)]
```
### HGET: trả về value dựa trên key
```
  HGET <id> <key>
```
### HGETALL: trả về tất cả các cặp key-val
```
  HGETALL <id>
```
### HEXISTS: Kiểm tra xem key đã tồn tại trong hash chưa.
```
  HEXISTS <id> <key>
```
### DEL: Xóa toàn bộ key-value có trong hash
```
  DEL <id>
```
### HDEL: Xóa MỘT cặp key-value
```
  HDEL <id> <key>
```
### HINCRBY/HINCRBYFLOAT: Tăng/giảm một lượng với các value(int/float) có kiểu dữ liệu là int, float
```
  HINCRBY/HINCRBYFLOAT <id> <key> <x>
```
### HSTRLEN: Trả về độ dài string của value.
```
  HSTRLEN <id> <key>
```
### HKEYS: Liệt kê các key được định nghĩa trong hash
```
  HKEYS <id>
```
### HVALS: Liệt kê các value được định nghĩa trong hash
```
  HVALS <id>
```

## 3. Set
- Các dữ liệu trong set là duy nhất.
- Có thể thực hiện các phép toán của set như union, intersection, difference.
### SADD: Thêm 1 phần tử vào set.
```
  SADD <id> <value>
```
### SMEMBERS: Trả về tất cả phần tử có trong set.
```
  SMEMBER <id>
```
### SUNION: Trả về kết quả của phép hợp giữa các set khác nhau.
```
  SUNION [(các_set)]
```
### SINTER: Trả về kết quả của phép giao giữa các set khác nhau.
```
  SINTER [(các_set)]
```
### SDIFF: Trả về kết quả của phép trừ giữa các set khác nhau (các phần tử tồn tại ở set_1 nhưng kh tồn tại ở các set còn lại).
```
  SDIFF [(các_set)]
```
### SINTERSTORE: Lưu kết quả của các phép tập hợp.
```
  SINTERSTORE <cur_set> [các_set]
```
### SISMEMBER: Trả về 1 nếu phần tử có trong set và ngược lại
```
  SISMEMBER <id> <value>
```
### SMISMEMBER: Giống SISMEMBER nhưng kiểm tra nhiều phần tử trong 1 request;
```
  SMISMEMBER <id> [multiple_val]
```
### SCARD: Trả về SL phần tử có trong set.
```
  SCARD <id>
```
### SREM: Xóa phần tử có trong set.
```
  SREM <id> <val>
```
### SSCAN: Duyệt qua các phần tử với offset cho trước
```
  SSCAN <id> <offset> (option LIMIT <số_PT>)
```

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
## 5. Hyperloglog: đếm số lượng phần tử duy nhất được thêm vào hyperloglog với dung lượng thấp và tính chính xác tương đối.
### PFADD: thêm một phần tử.
```
  PFADD <id> <value>
```
### PFCOUNT: Đếm số lượng phần tử có trong id
```
  PFCOUNT <id>
```

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
