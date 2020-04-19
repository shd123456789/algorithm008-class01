<?php
// 355. 设计推特
class Twitter 
{
    private $followList = [];
    private $twitterContent = [];
    private $sort = 1; // 用时间戳导致插入时间相同排序会出问题

    function postTweet($userId, $tweetId) 
    {   
        $this->twitterContent[$userId][] = ['id' => $tweetId, 'sort' => $this->sort++];
    }
  
    function getNewsFeed($userId) 
    {   
        $myFollow = $this->getMyfollow($userId);
        $myFollow[$userId] = 1; 

        $twitterList = [];
        foreach ($myFollow as $key => $value) {
            $count = count($this->twitterContent[$key]) - 1;
            $j = 0;
            for($i = $count; $i >=0; $i--){    
                if ($j > 9) break; // 每个用户取前10条
                $twitterList[] = $this->twitterContent[$key][$i];
                $j++;
            }
        }

        array_multisort(array_column($twitterList,'sort'),SORT_DESC,$twitterList); //排序

        $tweetIds = [];
        foreach ($twitterList as $key => $value) {
            if ($key > 9) break;
            $tweetIds[] = $value['id'];
        }

        return $tweetIds;
    }

    function follow($followerId, $followeeId) 
    {
        $this->followList[$followerId][$followeeId] = 1;
    }

    function unfollow($followerId, $followeeId) 
    {
        unset($this->followList[$followerId][$followeeId]);
    }

    private function getMyfollow($followerId)
    {   
        if (empty($this->followList[$followerId])) return [];

        return $this->followList[$followerId]; 
    }

}