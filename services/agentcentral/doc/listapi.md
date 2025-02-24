# List APIs

## 需求 (V1)

1. 可以分页
2. 可以用简单的条件filter

## 优化

### filter

|Type|Parameter Name| Description |
|:---:|:---:|:---:|
| string | prefix | 匹配id / name的前缀 |
| int64 | createAfter | 匹配创建时间在createAfter之后或者等于createAfter的元素 |
| string | idAfter | 匹配primaryKey的字典序在idAfter之后或等于idAfter的元素 |
| int | limit | 最多取出limit个元素 |

## 需求 (V2)

TODO.