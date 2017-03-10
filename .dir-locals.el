((nil . ((eval .
  (and buffer-file-name
       (let ((dir (car (dir-locals-find-file (buffer-file-name)))))
         (visit-tags-table (expand-file-name "TAGS" dir) t)))))))
