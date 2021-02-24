# albumservice


AlbumPath:
    -album1
        -pic1-mini.jpg
        -pic1-max.jpg
        -pic1-org.jpg
        -album.json
            {
                Name  :string=album1
	            Cover string
	            Date  string
            }
    -album2
    
    
统一请求进入Controller的func，在此func 中使用反射根据路由map到具体func，并过滤GET或POST，并根据反射到的func 参数，读Request body。
