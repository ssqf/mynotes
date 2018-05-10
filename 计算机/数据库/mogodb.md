

1. 分组
db.unicorns.aggregate([{$group:{_id:'$gender', total: {$sum:1}}}])